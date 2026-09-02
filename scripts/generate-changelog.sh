#!/usr/bin/env bash
set -euo pipefail

PREV_TAG="$1"   # e.g. v1.2.2, or "" if this is the first release
CURR_TAG="$2"   # e.g. v1.2.3
REPO="yashranjan1/relay"

# Resolve date range from tag commits
if [ -n "$PREV_TAG" ]; then
  SINCE=$(git log -1 --format=%aI "$PREV_TAG")
else
  SINCE="1970-01-01T00:00:00Z"
fi
UNTIL=$(git log -1 --format=%aI "$CURR_TAG")

echo "## Changelog"
echo ""

# Merge-strategy-agnostic: find merged PRs into base branch in this window
PR_NUMBERS=$(gh pr list \
  --repo "$REPO" \
  --state merged \
  --base main \
  --search "merged:${SINCE}..${UNTIL}" \
  --json number -q '.[].number')

FOUND_ANY=false

for PR in $PR_NUMBERS; do
  ISSUE_NUMBERS=$(gh pr view "$PR" --repo "$REPO" \
    --json closingIssuesReferences \
    -q '.closingIssuesReferences[].number' 2>/dev/null || true)

  if [ -z "$ISSUE_NUMBERS" ]; then
    continue
  fi

  for ISSUE in $ISSUE_NUMBERS; do
    BODY=$(gh issue view "$ISSUE" --repo "$REPO" --json body -q '.body')

    # Issue-form fields render as "### <Label>" headings
    CRITERIA=$(echo "$BODY" | awk '/^### Satisfaction Criteria/{flag=1; next} /^### /{flag=0} flag')
    CRITERIA=$(echo "$CRITERIA" | sed '/^\s*$/d')

    if [ -n "$CRITERIA" ]; then
      FOUND_ANY=true
      TITLE=$(gh issue view "$ISSUE" --repo "$REPO" --json title -q '.title')
      echo "### $TITLE (#$ISSUE, via PR #$PR)"
      echo "$CRITERIA"
      echo ""
    fi
  done
done

if [ "$FOUND_ANY" = false ]; then
  echo "_No satisfaction criteria recorded for this release._"
fi
